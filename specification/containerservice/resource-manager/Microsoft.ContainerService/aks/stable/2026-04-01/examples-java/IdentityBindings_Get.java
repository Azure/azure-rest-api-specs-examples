
/**
 * Samples for IdentityBindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/IdentityBindings_Get.json
     */
    /**
     * Sample code: Get Identity Binding.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getIdentityBinding(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getIdentityBindings().getWithResponse("rg1", "clustername1", "identitybinding1",
            com.azure.core.util.Context.NONE);
    }
}
