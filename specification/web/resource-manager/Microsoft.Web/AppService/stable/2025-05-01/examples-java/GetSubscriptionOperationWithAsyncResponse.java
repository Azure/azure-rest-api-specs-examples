
/**
 * Samples for Global GetSubscriptionOperationWithAsyncResponse.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSubscriptionOperationWithAsyncResponse.json
     */
    /**
     * Sample code: Gets an operation in a subscription and given region.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getsAnOperationInASubscriptionAndGivenRegion(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getGlobals().getSubscriptionOperationWithAsyncResponseWithResponse("West US",
            "34adfa4f-cedf-4dc0-ba29-b6d1a69ab5d5", com.azure.core.util.Context.NONE);
    }
}
