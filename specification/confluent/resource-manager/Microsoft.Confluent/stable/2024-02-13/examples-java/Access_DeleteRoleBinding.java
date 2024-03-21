
/**
 * Samples for Access DeleteRoleBinding.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_DeleteRoleBinding.
     * json
     */
    /**
     * Sample code: Access_DeleteRoleBinding.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessDeleteRoleBinding(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().deleteRoleBindingWithResponse("myResourceGroup", "myOrganization", "dlz-f3a90de",
            com.azure.core.util.Context.NONE);
    }
}
