
ª
noteapp.protoproto"b
User
id (Rid
email (	Remail
name (	Rname 
profile_pic (	Rprofile_pic"_
CreateUserRequest
email (	Remail
name (	Rname 
profile_pic (	Rprofile_pic"-
GetUserByEmailRequest
email (	Remail"_
UpdateUserRequest
email (	Remail
name (	Rname 
profile_pic (	Rprofile_pic")
DeleteUserRequest
email (	Remail".
DeleteUserResponse
message (	Rmessage"/
UserResponse
user (2.proto.UserRuser"`
Note
id (Rid
user_id (Ruser_id
title (	Rtitle
content (	Rcontent"]
CreateNoteRequest
user_id (Ruser_id
title (	Rtitle
content (	Rcontent"$
GetNoteByIDRequest
id (Rid"3
GetNotesByUserIDRequest
user_id (Ruser_id"=
GetNotesByUserIDResponse!
notes (2.proto.NoteRnotes"S
UpdateNoteRequest
id (Rid
title (	Rtitle
content (	Rcontent"#
DeleteNoteRequest
id (Rid".
DeleteNoteResponse
message (	Rmessage"/
NoteResponse
note (2.proto.NoteRnote2è
UserService;

CreateUser.proto.CreateUserRequest.proto.UserResponseC
GetUserByEmail.proto.GetUserByEmailRequest.proto.UserResponse;

UpdateUser.proto.UpdateUserRequest.proto.UserResponseA

DeleteUser.proto.DeleteUserRequest.proto.DeleteUserResponse2ﬁ
NoteService;

CreateNote.proto.CreateNoteRequest.proto.NoteResponse=
GetNoteByID.proto.GetNoteByIDRequest.proto.NoteResponseS
GetNotesByUserID.proto.GetNotesByUserIDRequest.proto.GetNotesByUserIDResponse;

UpdateNote.proto.UpdateNoteRequest.proto.NoteResponseA

DeleteNote.proto.DeleteNoteRequest.proto.DeleteNoteResponseBZproto/generatedbproto3