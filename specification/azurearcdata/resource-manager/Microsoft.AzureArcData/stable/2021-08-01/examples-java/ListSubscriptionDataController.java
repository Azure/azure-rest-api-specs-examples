import com.azure.core.util.Context;

/** Samples for DataControllers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListSubscriptionDataController.json
     */
    /**
     * Sample code: Gets all dataControllers in a subscription.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllDataControllersInASubscription(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.dataControllers().list(Context.NONE);
    }
}
