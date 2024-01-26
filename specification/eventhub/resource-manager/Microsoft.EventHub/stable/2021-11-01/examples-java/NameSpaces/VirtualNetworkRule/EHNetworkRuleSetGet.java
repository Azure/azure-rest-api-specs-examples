
import com.azure.core.util.Context;

/** Samples for Namespaces GetNetworkRuleSet. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * VirtualNetworkRule/EHNetworkRuleSetGet.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceNetworkRuleSetGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().getNetworkRuleSetWithResponse("ResourceGroup",
            "sdk-Namespace-6019", Context.NONE);
    }
}
