
import com.azure.resourcemanager.cognitiveservices.fluent.models.CapabilityHostInner;
import com.azure.resourcemanager.cognitiveservices.models.CapabilityHostProperties;

/**
 * Samples for ProjectCapabilityHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ProjectCapabilityHost/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Project CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createOrUpdateProjectCapabilityHost(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCapabilityHosts().createOrUpdate("test-rg", "account-1", "project-1", "capabilityHostName",
            new CapabilityHostInner().withProperties(new CapabilityHostProperties().withCustomerSubnet(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet")),
            com.azure.core.util.Context.NONE);
    }
}
