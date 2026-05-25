
/**
 * Samples for IdentityBindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/IdentityBindings_Delete.json
     */
    /**
     * Sample code: Delete Identity Binding.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteIdentityBinding(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getIdentityBindings().delete("rg1", "clustername1", "identitybinding1",
            com.azure.core.util.Context.NONE);
    }
}
