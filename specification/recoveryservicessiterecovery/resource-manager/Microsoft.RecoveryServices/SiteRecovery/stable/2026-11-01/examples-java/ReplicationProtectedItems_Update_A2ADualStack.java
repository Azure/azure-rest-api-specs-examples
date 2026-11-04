
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AUpdateReplicationProtectedItemInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.IpConfigInputDetails;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReplicationProtectedItem;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateReplicationProtectedItemInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VMNicInputDetails;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectedItems Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Update_A2ADualStack.json
     */
    /**
     * Sample code: Updates the recovery ip-configs of a dual-stack (IPv4 and IPv6) A2A protected item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updatesTheRecoveryIpConfigsOfADualStackIPv4AndIPv6A2AProtectedItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        ReplicationProtectedItem resource
            = manager.replicationProtectedItems().getWithResponse("priyanprg", "priyanponeboxvault", "CentralUSCanSite",
                "CentralUSCancloud", "a2aDualStackVm", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new UpdateReplicationProtectedItemInputProperties()
            .withSelectedRecoveryAzureNetworkId(
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/virtualNetworks/dualstack-recovery-vnet")
            .withVmNics(Arrays.asList(new VMNicInputDetails().withNicId(
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/networkInterfaces/a2aDualStackVm-nic")
                .withIpConfigs(Arrays.asList(
                    new IpConfigInputDetails().withIpConfigName("ipconfig1").withIsPrimary(true)
                        .withIsSeletedForFailover(true).withRecoverySubnetName("dualstack-subnet")
                        .withRecoveryStaticIpAddress("10.0.2.46"),
                    new IpConfigInputDetails().withIpConfigName("ipconfig-v6").withIsPrimary(false)
                        .withIsSeletedForFailover(true).withRecoverySubnetName("dualstack-subnet")
                        .withRecoveryStaticIpAddress("2001:db8:1::4")))
                .withSelectionType("SelectedByUser")))
            .withProviderSpecificDetails(new A2AUpdateReplicationProtectedItemInput().withRecoveryResourceGroupId(
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg")))
            .apply();
    }
}
