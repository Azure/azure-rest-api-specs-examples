
/**
 * Samples for TrustedAccessRoleBindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/TrustedAccessRoleBindings_Delete.json
     */
    /**
     * Sample code: Delete a trusted access role binding.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteATrustedAccessRoleBinding(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getTrustedAccessRoleBindings().delete("rg1", "clustername1", "binding1",
            com.azure.core.util.Context.NONE);
    }
}
