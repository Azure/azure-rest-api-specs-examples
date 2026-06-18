
/**
 * Samples for PrivateAccesses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_Delete.json
     */
    /**
     * Sample code: Delete a private access resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAPrivateAccessResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().delete("myResourceGroup", "myPrivateAccess", com.azure.core.util.Context.NONE);
    }
}
