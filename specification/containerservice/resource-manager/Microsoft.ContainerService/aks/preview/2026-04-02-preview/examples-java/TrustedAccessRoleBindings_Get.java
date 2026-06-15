
/**
 * Samples for TrustedAccessRoleBindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/TrustedAccessRoleBindings_Get.json
     */
    /**
     * Sample code: Get a trusted access role binding.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getATrustedAccessRoleBinding(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getTrustedAccessRoleBindings().getWithResponse("rg1", "clustername1", "binding1",
            com.azure.core.util.Context.NONE);
    }
}
