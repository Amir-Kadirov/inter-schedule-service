-- Create the EnglishLevel enum type
CREATE TYPE "EnglishLevel" AS ENUM (
  'Beginner',
  'Elementary',
  'PreIntermediate',
  'Intermediate',
  'UpperIntermediate',
  'Ielts'
);

CREATE TABLE IF NOT EXISTS "Student" (
  "ID" uuid PRIMARY KEY,
  "FullName" varchar(100),
  "Phone" varchar(15),
  "Email" varchar,
  "PaidSum" int,
  "CourseCount" int,
  "Password" varchar(255),
  "GroupID" uuid,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Group_Table" (
  "StudentID" uuid,
  "GroupID" uuid
);

CREATE TABLE IF NOT EXISTS "Group" (
  "ID" uuid PRIMARY KEY,
  "TeacherID" uuid,
  "SupportTeacherID" uuid,
  "BranchID" uuid,
  "Type" "EnglishLevel",
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Group_Many" (
  "GroupID" uuid REFERENCES "Group" ("ID"),
  "ScheduleID" uuid REFERENCES "Schedule" ("ID"),
  "JournalID" uuid REFERENCES "Journal" ("ID")
);

CREATE TABLE IF NOT EXISTS "Schedule" (
  "ID" uuid PRIMARY KEY,
  "StartTime" timestamp,
  "EndTime" timestamp,
  "LessonID" uuid,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Schedule_Lesson" (
  "LessonID" uuid,
  "ScheduleID" uuid
);

CREATE TABLE IF NOT EXISTS "Task" (
  "ID" uuid PRIMARY KEY,
  "LessonID" uuid,
  "Label" varchar,
  "Deadline" varchar,
  "Score" int,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Create the Lesson_Task table
CREATE TABLE IF NOT EXISTS "Lesson_Task" (
  "LessonID" uuid,
  "TaskID" uuid
);

-- Create the Lesson table
CREATE TABLE IF NOT EXISTS "Lesson" (
  "ID" uuid PRIMARY KEY,
  "Name" varchar
);

-- Create the Journal table
CREATE TABLE IF NOT EXISTS "Journal" (
  "ID" uuid PRIMARY KEY,
  "From" timestamp,
  "To" timestamp,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Create the Register_Event table
CREATE TABLE IF NOT EXISTS "Register_Event" (
  "EventID" uuid,
  "StudentID" uuid,
  "BranchID" uuid
);

-- Create the Event table
CREATE TABLE IF NOT EXISTS "Event" (
  "ID" uuid PRIMARY KEY,
  "Topic" varchar,
  "StartTime" timestamp,
  "Day" date,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Add foreign keys to the Student table
ALTER TABLE "Student" ADD FOREIGN KEY ("GroupID") REFERENCES "Group" ("ID");

ALTER TABLE "Student" ADD COLUMN "TotalSum" int;
ALTER TABLE "Student" ADD COLUMN "Score" int;

ALTER TABLE "Lesson" ADD COLUMN "created_at" timestamp DEFAULT NOW();
ALTER TABLE "Lesson" ADD COLUMN "deleted_at" timestamp;
ALTER TABLE "Lesson" ADD COLUMN "updated_at" timestamp;

ALTER TABLE "Register_Event" ADD FOREIGN KEY ("EventID") REFERENCES "Event" ("ID");
ALTER TABLE "Register_Event" ADD FOREIGN KEY ("StudentID") REFERENCES "Student" ("ID");

-- Add foreign keys to the Group_Table table
ALTER TABLE "Group_Table" ADD FOREIGN KEY ("StudentID") REFERENCES "Student" ("ID");
ALTER TABLE "Group_Table" ADD FOREIGN KEY ("GroupID") REFERENCES "Group" ("ID");

-- Add foreign keys to the Schedule table
ALTER TABLE "Schedule" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");

-- Add foreign keys to the Schedule_Lesson table
ALTER TABLE "Schedule_Lesson" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");
ALTER TABLE "Schedule_Lesson" ADD FOREIGN KEY ("ScheduleID") REFERENCES "Schedule" ("ID");

-- Add foreign keys to the Task table
ALTER TABLE "Task" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");
ALTER TABLE "Task"
ALTER COLUMN "Deadline" TYPE DATE USING "Deadline"::DATE;
ALTER TABLE "Task" ADD COLUMN "Status" BOOLEAN;

-- Add foreign keys to the Lesson_Task table
ALTER TABLE "Lesson_Task" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");
ALTER TABLE "Lesson_Task" ADD FOREIGN KEY ("TaskID") REFERENCES "Task" ("ID");

ALTER TABLE "Task"
ALTER COLUMN "Status"
SET DEFAULT FALSE;

ALTER TABLE "Task"
ALTER COLUMN "Score"
SET DEFAULT 0;

ALTER TABLE "Student"
ALTER COLUMN "PaidSum"
SET DEFAULT 0;

ALTER TABLE "Student"
ALTER COLUMN "CourseCount"
SET DEFAULT 0;

ALTER TABLE "Student"
ALTER COLUMN "TotalSum"
SET DEFAULT 0;

CREATE OR REPLACE FUNCTION after_insert_student()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO "Group_Table" ("StudentID", "GroupID")
    VALUES (NEW."ID", NEW."GroupID");
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for the Student table
CREATE TRIGGER after_insert_student_trigger
AFTER INSERT ON "Student"
FOR EACH ROW
EXECUTE FUNCTION after_insert_student();

-- Create after insert trigger function for the Task table
CREATE OR REPLACE FUNCTION after_insert_task()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO "Lesson_Task" ("LessonID", "TaskID")
    VALUES (NEW."LessonID", NEW."ID");
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for the Task table
CREATE TRIGGER after_insert_task_trigger
AFTER INSERT ON "Task"
FOR EACH ROW
EXECUTE FUNCTION after_insert_task();

-- Create after insert trigger function for the Schedule table
CREATE OR REPLACE FUNCTION after_insert_schedule()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO "Schedule_Lesson" ("LessonID", "ScheduleID")
    VALUES (NEW."LessonID", NEW."ID");
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for the Schedule table
CREATE TRIGGER after_insert_schedule_trigger
AFTER INSERT ON "Schedule"
FOR EACH ROW
EXECUTE FUNCTION after_insert_schedule();
