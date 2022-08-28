import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs/SBAliasList.json
     */
    /**
     * Sample code: SBAliasList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sBAliasList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .serviceBusNamespaces()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .list("ardsouzatestRG", "sdk-Namespace-8860", Context.NONE);
    }
}
