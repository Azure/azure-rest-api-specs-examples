import com.azure.core.util.Context;

/** Samples for ProvisionedClustersOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListProvisionedClusterBySubscription.json
     */
    /**
     * Sample code: ListProvisionedClusterBySubscription.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listProvisionedClusterBySubscription(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.provisionedClustersOperations().list(Context.NONE);
    }
}
