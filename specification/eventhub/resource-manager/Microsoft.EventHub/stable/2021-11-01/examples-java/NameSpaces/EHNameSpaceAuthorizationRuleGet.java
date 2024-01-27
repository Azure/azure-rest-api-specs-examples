
import com.azure.core.util.Context;

/** Samples for Namespaces GetAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * EHNameSpaceAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().getAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-2702", "sdk-Authrules-1746", Context.NONE);
    }
}
