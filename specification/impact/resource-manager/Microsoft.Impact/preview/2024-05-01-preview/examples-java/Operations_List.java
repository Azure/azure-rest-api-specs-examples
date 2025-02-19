
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Operations_List.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void operationsList(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
