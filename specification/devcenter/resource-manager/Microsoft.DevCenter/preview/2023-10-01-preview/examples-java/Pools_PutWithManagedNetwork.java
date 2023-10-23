import com.azure.resourcemanager.devcenter.models.LicenseType;
import com.azure.resourcemanager.devcenter.models.LocalAdminStatus;
import com.azure.resourcemanager.devcenter.models.SingleSignOnStatus;
import com.azure.resourcemanager.devcenter.models.StopOnDisconnectConfiguration;
import com.azure.resourcemanager.devcenter.models.StopOnDisconnectEnableStatus;
import com.azure.resourcemanager.devcenter.models.VirtualNetworkType;
import java.util.Arrays;

/** Samples for Pools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_PutWithManagedNetwork.json
     */
    /**
     * Sample code: Pools_CreateOrUpdateWithManagedNetwork.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsCreateOrUpdateWithManagedNetwork(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .pools()
            .define("DevPool")
            .withRegion("centralus")
            .withExistingProject("rg1", "DevProject")
            .withDevBoxDefinitionName("WebDevBox")
            .withNetworkConnectionName("managedNetwork")
            .withLicenseType(LicenseType.WINDOWS_CLIENT)
            .withLocalAdministrator(LocalAdminStatus.ENABLED)
            .withStopOnDisconnect(
                new StopOnDisconnectConfiguration()
                    .withStatus(StopOnDisconnectEnableStatus.ENABLED)
                    .withGracePeriodMinutes(60))
            .withSingleSignOnStatus(SingleSignOnStatus.DISABLED)
            .withDisplayName("Developer Pool")
            .withVirtualNetworkType(VirtualNetworkType.MANAGED)
            .withManagedVirtualNetworkRegions(Arrays.asList("centralus"))
            .create();
    }
}
