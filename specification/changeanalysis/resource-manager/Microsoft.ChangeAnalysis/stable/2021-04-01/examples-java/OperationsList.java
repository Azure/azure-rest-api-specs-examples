/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/OperationsList.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AzureChangeAnalysisManager.
     */
    public static void operationsList(com.azure.resourcemanager.changeanalysis.AzureChangeAnalysisManager manager) {
        manager.operations().list(null, com.azure.core.util.Context.NONE);
    }
}
