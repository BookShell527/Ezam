import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';
import { School } from './school';
import { Teacher } from './teacher';
export type ClassDocument = HydratedDocument<Class>;

@Schema()
export class Class {
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'schools' })
  school: School;
  @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'teachers' })
  teacher: Teacher;

  @Prop({ required: true })
  name: string;
  @Prop({ required: true })
  semester: number[];
  @Prop({ required: true, default: (new Date().getFullYear()).toString() + "/" + ((new Date().getFullYear() + 1).toString()) })
  academic_year: string;
}

export const ClassSchema = SchemaFactory.createForClass(Class);
