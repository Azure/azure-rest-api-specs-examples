import com.azure.core.util.Context;

/** Samples for MobileNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/MobileNetworkListByResourceGroup.json
     */
    /**
     * Sample code: List mobile networks in resource group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listMobileNetworksInResourceGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().listByResourceGroup("rg1", Context.NONE);
    }
}
