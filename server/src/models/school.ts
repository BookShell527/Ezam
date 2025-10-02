import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type SchoolDocument = HydratedDocument<School>;

type EducationalLevel = "Elementary School" | "Middle School" | "High School" | "University & Above";

@Schema()
export class School {
  @Prop({ required: true })
  name: string;
  @Prop({ required: true })
  address: string;
  @Prop({ required: true })
  telp_no: string;
  @Prop({ required: true })
  school_email: string;
  @Prop({ required: true })
  educational_level: EducationalLevel;
}

export const SchoolSchema = SchemaFactory.createForClass(School);
