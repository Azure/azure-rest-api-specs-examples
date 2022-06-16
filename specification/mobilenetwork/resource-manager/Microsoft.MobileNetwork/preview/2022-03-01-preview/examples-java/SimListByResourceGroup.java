import com.azure.core.util.Context;

/** Samples for Sims ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimListByResourceGroup.json
     */
    /**
     * Sample code: List sims in a resource group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSimsInAResourceGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().listByResourceGroup("rg1", Context.NONE);
    }
}
