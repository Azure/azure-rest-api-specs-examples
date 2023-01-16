/** Samples for HanaInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_List.json
     */
    /**
     * Sample code: List all HANA instances in a subscription.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void listAllHANAInstancesInASubscription(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.hanaInstances().list(com.azure.core.util.Context.NONE);
    }
}
