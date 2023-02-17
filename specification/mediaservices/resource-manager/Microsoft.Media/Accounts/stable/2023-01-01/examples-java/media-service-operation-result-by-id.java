/** Samples for MediaServicesOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/media-service-operation-result-by-id.json
     */
    /**
     * Sample code: Get status of asynchronous operation.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperation(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaServicesOperationResults()
            .getWithResponse("westus", "6FBA62C4-99B5-4FF8-9826-FC4744A8864F", com.azure.core.util.Context.NONE);
    }
}
