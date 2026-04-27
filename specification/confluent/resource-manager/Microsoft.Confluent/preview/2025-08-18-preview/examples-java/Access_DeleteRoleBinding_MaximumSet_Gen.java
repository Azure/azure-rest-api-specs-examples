
/**
 * Samples for Access DeleteRoleBinding.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_DeleteRoleBinding_MaximumSet_Gen.json
     */
    /**
     * Sample code: Access_DeleteRoleBinding_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessDeleteRoleBindingMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().deleteRoleBindingWithResponse("rgconfluent", "aeqwsawfoagclmfwwaw",
            "ucuqvcuiwmoreczccknufbhrwyp", com.azure.core.util.Context.NONE);
    }
}
