import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument, Document } from 'mongoose';
import { Class } from './class';
import { Teacher } from './teacher';
export type ExamDocument = HydratedDocument<Exam>;


@Schema({ _id: false })
export class Question extends Document {
  @Prop({ required: true })
  index: number;
  @Prop({ required: true })
  question: string;
  @Prop({ required: true })
  answer: string;
  @Prop({ default: [] })
  options: string[];
  @Prop({ default: -1 })
  point: number;
}
export const QuestionSchema = SchemaFactory.createForClass(Question);

@Schema()
export class Exam {
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'classes' })
  class: Class;
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'teachers' })
  teacher: Teacher;

  @Prop({ required: true })
  name: string;
  @Prop({ required: true })
  subject: string;
  @Prop({ default: false })
  shuffle_question: boolean;
  @Prop({ default: false })
  shuffle_answer: boolean;
  @Prop({ required: true })
  start_time: Date;
  @Prop({ required: true })
  end_time: Date;
  @Prop({ required: true })
  duration: number;
  @Prop({ default: 100 })
  max_score: number;

  @Prop({ type: [QuestionSchema], default: [] })
  question: Question[];
}

export const ExamSchema = SchemaFactory.createForClass(Exam);
