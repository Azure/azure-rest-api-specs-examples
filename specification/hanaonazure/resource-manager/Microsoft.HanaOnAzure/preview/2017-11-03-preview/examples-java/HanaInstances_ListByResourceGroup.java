/** Samples for HanaInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_ListByResourceGroup.json
     */
    /**
     * Sample code: List all HANA instances in a resource group.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void listAllHANAInstancesInAResourceGroup(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.hanaInstances().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
