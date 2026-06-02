
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/MsiOperationsList.json
     */
    /**
     * Sample code: MsiOperationsList.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void msiOperationsList(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
