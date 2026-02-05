
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/OperationsList.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void operationsList(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.operations().list(null, com.azure.core.util.Context.NONE);
    }
}
