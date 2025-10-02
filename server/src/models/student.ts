import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument, Document } from 'mongoose';
import type { Date } from 'mongoose';
import { School } from './school';
import { Class } from './class';
import { Exam } from './exam';
export type StudentDocument = HydratedDocument<Student>;


interface Answer {
  index: string;
  answer: string;
}

@Schema({ _id: false })
export class TakenExam extends Document {
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'exams' })
  exam: Exam;

  @Prop({ required: true })
  student_start_time: Date;
  @Prop()
  student_end_time: Date;

  @Prop({ default: [] })
  answers: Answer[];
}
export const TakenExamSchema = SchemaFactory.createForClass(TakenExam);

@Schema()
export class Student {
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'schools' })
  school: School;
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'classes' })
  class: Class;
  @Prop({ type: [TakenExamSchema], default: [] })
  taken_exam: TakenExam[];

  @Prop({ required: true })
  student_id: string;
  @Prop({ required: true })
  name: string;
  @Prop({ required: true })
  email: string;
  @Prop({ required: true })
  date_of_birth: Date;
  @Prop({ required: true })
  gender: "Male" | "Female";
  @Prop({ required: true })
  address: string;
  @Prop({ required: true })
  enrollment_date: Date;
  @Prop({ required: true })
  status: "Active" | "Graduated" | "Suspended" | "Kicked Out";
}

export const StudentSchema = SchemaFactory.createForClass(Student);
