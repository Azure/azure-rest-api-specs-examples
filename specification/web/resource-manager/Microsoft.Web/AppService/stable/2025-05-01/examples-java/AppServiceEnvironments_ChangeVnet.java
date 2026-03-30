
import com.azure.resourcemanager.appservice.models.VirtualNetworkProfile;

/**
 * Samples for AppServiceEnvironments ChangeVnet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ChangeVnet.json
     */
    /**
     * Sample code: Move an App Service Environment to a different VNET.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        moveAnAppServiceEnvironmentToADifferentVNET(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().changeVnet("test-rg", "test-ase",
            new VirtualNetworkProfile().withId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
            com.azure.core.util.Context.NONE);
    }
}
