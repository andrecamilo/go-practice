Best practices
    context.Background should be used only at the highest level, as the root of all derived contexts
    context.TODO should be used where not sure what to use or if the current function will be updated to use context in future
    context cancelations are advisory, the functions may take time to clean up and exit
    context.Value should be used very rarely, it should never be used to pass in optional parameters. This makes the API implicit and can introduce bugs. Instead, such values should be passed in as arguments.
    Don’t store contexts in a struct, pass them explicitly in functions, preferably, as the first argument.
    Never pass nil context, instead, use a TODO if you are not sure what to use.
    The Context struct does not have a cancel method because only the function that derives the context should cancel it.

References:
    The Context struct does not have a cancel method because only the function that derives the context should cancel it.