
import com.azure.resourcemanager.servicebus.fluent.models.ArmDisasterRecoveryInner;

/**
 * Samples for DisasterRecoveryConfigs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasCreate.json
     */
    /**
     * Sample code: SBAliasCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBAliasCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().createOrUpdateWithResponse("ardsouzatestRG",
            "sdk-Namespace-8860", "sdk-Namespace-8860", new ArmDisasterRecoveryInner()
                .withPartnerNamespace("sdk-Namespace-37").withAlternateName("alternameforAlias-Namespace-8860"),
            com.azure.core.util.Context.NONE);
    }
}
