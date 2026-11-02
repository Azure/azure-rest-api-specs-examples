
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void
        operationsListMaximumSet(com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
