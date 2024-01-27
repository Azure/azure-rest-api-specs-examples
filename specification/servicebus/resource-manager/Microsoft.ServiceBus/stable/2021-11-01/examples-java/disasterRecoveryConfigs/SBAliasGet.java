
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasGet.json
     */
    /**
     * Sample code: SBAliasGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sBAliasGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs()
            .getWithResponse("ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", Context.NONE);
    }
}
