
/**
 * Samples for Access DeleteRoleBinding.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_DeleteRoleBinding_MinimumSet_Gen.json
     */
    /**
     * Sample code: Access_DeleteRoleBinding_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessDeleteRoleBindingMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().deleteRoleBindingWithResponse("rgconfluent", "kxbwvfhsqesuaswozdiivwo", "dqlmrdp",
            com.azure.core.util.Context.NONE);
    }
}
