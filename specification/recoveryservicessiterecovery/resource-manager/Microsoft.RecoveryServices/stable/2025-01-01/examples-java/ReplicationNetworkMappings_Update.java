
import com.azure.resourcemanager.recoveryservicessiterecovery.models.NetworkMapping;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateNetworkMappingInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VmmToAzureUpdateNetworkMappingInput;

/**
 * Samples for ReplicationNetworkMappings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationNetworkMappings_Update.json
     */
    /**
     * Sample code: Updates network mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        updatesNetworkMapping(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        NetworkMapping resource = manager.replicationNetworkMappings()
            .getWithResponse("srcBvte2a14C27", "srce2avaultbvtaC27",
                "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac",
                "e2267b5c-2650-49bd-ab3f-d66aae694c06", "corpe2amap", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new UpdateNetworkMappingInputProperties()
            .withRecoveryFabricName("Microsoft Azure")
            .withRecoveryNetworkId(
                "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai2")
            .withFabricSpecificDetails(new VmmToAzureUpdateNetworkMappingInput())).apply();
    }
}
