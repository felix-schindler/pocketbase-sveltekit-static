import type { AuthRecord, RecordService } from 'pocketbase';
import type PocketBase from 'pocketbase';

export type TypedPocketBase = PocketBase & {
	collection(idOrName: 'users'): RecordService<User>;
};

export type User = AuthRecord & {
	name: string;
};
