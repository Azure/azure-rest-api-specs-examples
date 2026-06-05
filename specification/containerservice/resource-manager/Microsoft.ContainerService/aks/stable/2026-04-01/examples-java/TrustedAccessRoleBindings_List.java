
/**
 * Samples for TrustedAccessRoleBindings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/TrustedAccessRoleBindings_List.json
     */
    /**
     * Sample code: List trusted access role bindings.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listTrustedAccessRoleBindings(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getTrustedAccessRoleBindings().list("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
