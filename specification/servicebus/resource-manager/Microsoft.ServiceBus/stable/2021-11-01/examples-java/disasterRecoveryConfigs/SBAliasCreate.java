
import com.azure.resourcemanager.servicebus.fluent.models.ArmDisasterRecoveryInner;

/**
 * Samples for DisasterRecoveryConfigs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasCreate.json
     */
    /**
     * Sample code: SBAliasCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sBAliasCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs().createOrUpdateWithResponse(
            "ardsouzatestRG", "sdk-Namespace-8860", "sdk-Namespace-8860", new ArmDisasterRecoveryInner()
                .withPartnerNamespace("sdk-Namespace-37").withAlternateName("alternameforAlias-Namespace-8860"),
            com.azure.core.util.Context.NONE);
    }
}
