import com.azure.resourcemanager.recoveryservices.models.AlertsState;
import com.azure.resourcemanager.recoveryservices.models.AzureMonitorAlertSettings;
import com.azure.resourcemanager.recoveryservices.models.ClassicAlertSettings;
import com.azure.resourcemanager.recoveryservices.models.MonitoringSettings;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Vaults Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PATCHVault_WithMonitoringSettings.json
     */
    /**
     * Sample code: Update Vault With Monitoring Setting.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateVaultWithMonitoringSetting(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource =
            manager
                .vaults()
                .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("PatchKey", "fakeTokenPlaceholder"))
            .withProperties(
                new VaultProperties()
                    .withMonitoringSettings(
                        new MonitoringSettings()
                            .withAzureMonitorAlertSettings(
                                new AzureMonitorAlertSettings().withAlertsForAllJobFailures(AlertsState.ENABLED))
                            .withClassicAlertSettings(
                                new ClassicAlertSettings().withAlertsForCriticalOperations(AlertsState.DISABLED))))
            .apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
