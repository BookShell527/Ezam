import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';
import type { Date } from 'mongoose';
import { School } from './school';
export type TeacherDocument = HydratedDocument<Teacher>;

@Schema()
export class Teacher {
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'schools' })
  school: School;

  @Prop({ required: true })
  teacher_id: string;
  @Prop({ required: true })
  name: string;
  @Prop({ required: true })
  subject: string[];
  @Prop({ required: true })
  gender: "Male" | "Female";
  @Prop({ required: true })
  date_of_birth: Date;
  @Prop({ required: true })
  telp_number: string;
  @Prop({ required: true })
  address: string;
  @Prop({ required: true })
  email: string;
}

export const TeacherSchema = SchemaFactory.createForClass(Teacher);
