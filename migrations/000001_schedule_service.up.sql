-- Create the EnglishLevel enum type
CREATE TYPE "EnglishLevel" AS ENUM (
  'Beginner',
  'Elementary',
  'PreIntermediate',
  'Intermediate',
  'UpperIntermediate',
  'Ielts'
);

-- Create the Student table
CREATE TABLE "Student" (
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

-- Create the Group_Table table
CREATE TABLE "Group_Table" (
  "StudentID" uuid,
  "GroupID" uuid
);

-- Create the Group table
CREATE TABLE "Group" (
  "ID" uuid PRIMARY KEY,
  "TeacherID" uuid,
  "SupportTeacherID" uuid,
  "JournalID" uuid,
  "BranchID" uuid,
  "Type" "EnglishLevel",
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Create the Schedule table
CREATE TABLE "Schedule" (
  "ID" uuid PRIMARY KEY,
  "StartTime" timestamp,
  "EndTime" timestamp,
  "JournalID" uuid,
  "LessonID" uuid,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Create the Schedule_Lesson table
CREATE TABLE "Schedule_Lesson" (
  "LessonID" uuid,
  "ScheduleID" uuid
);

-- Create the Task table
CREATE TABLE "Task" (
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
CREATE TABLE "Lesson_Task" (
  "LessonID" uuid,
  "TaskID" uuid
);

-- Create the Lesson table
CREATE TABLE "Lesson" (
  "ID" uuid PRIMARY KEY,
  "Name" varchar
);

-- Create the Journal table
CREATE TABLE "Journal" (
  "ID" uuid PRIMARY KEY,
  "From" timestamp,
  "To" timestamp,
  "ScheduleID" uuid,
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- Create the Register_Event table
CREATE TABLE "Register_Event" (
  "EventID" uuid,
  "StudentID" uuid,
  "BranchID" uuid
);

-- Create the Event table
CREATE TABLE "Event" (
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

-- Add foreign keys to the Register_Event table
ALTER TABLE "Register_Event" ADD FOREIGN KEY ("EventID") REFERENCES "Event" ("ID");
ALTER TABLE "Register_Event" ADD FOREIGN KEY ("StudentID") REFERENCES "Student" ("ID");

-- Add foreign keys to the Group_Table table
ALTER TABLE "Group_Table" ADD FOREIGN KEY ("StudentID") REFERENCES "Student" ("ID");
ALTER TABLE "Group_Table" ADD FOREIGN KEY ("GroupID") REFERENCES "Group" ("ID");

-- Add foreign keys to the Group table
ALTER TABLE "Group" ADD FOREIGN KEY ("JournalID") REFERENCES "Journal" ("ID");

-- Add foreign keys to the Schedule table
ALTER TABLE "Schedule" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");

-- Add foreign keys to the Schedule_Lesson table
ALTER TABLE "Schedule_Lesson" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");
ALTER TABLE "Schedule_Lesson" ADD FOREIGN KEY ("ScheduleID") REFERENCES "Schedule" ("ID");

-- Add foreign keys to the Task table
ALTER TABLE "Task" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");

-- Add foreign keys to the Lesson_Task table
ALTER TABLE "Lesson_Task" ADD FOREIGN KEY ("LessonID") REFERENCES "Lesson" ("ID");
ALTER TABLE "Lesson_Task" ADD FOREIGN KEY ("TaskID") REFERENCES "Task" ("ID");

-- Add foreign keys to the Journal table
ALTER TABLE "Journal" ADD FOREIGN KEY ("ScheduleID") REFERENCES "Schedule" ("ID");

-- Create after insert trigger function for the Student table
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
