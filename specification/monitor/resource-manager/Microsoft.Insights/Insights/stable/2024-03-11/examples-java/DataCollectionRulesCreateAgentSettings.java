
import com.azure.resourcemanager.monitor.fluent.models.DataCollectionRuleResourceInner;
import com.azure.resourcemanager.monitor.models.AgentSetting;
import com.azure.resourcemanager.monitor.models.DataCollectionRuleAgentSettings;
import com.azure.resourcemanager.monitor.models.KnownAgentSettingName;
import com.azure.resourcemanager.monitor.models.KnownDataCollectionRuleResourceKind;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DataCollectionRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateAgentSettings.json
     */
    /**
     * Sample code: Create or update an agent settings configuration.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        createOrUpdateAnAgentSettingsConfiguration(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().createWithResponse("myResourceGroup", "myCollectionRule",
            new DataCollectionRuleResourceInner().withLocation("eastus")
                .withKind(KnownDataCollectionRuleResourceKind.fromString("AgentSettings"))
                .withDescription("An agent settings configuration")
                .withAgentSettings(new DataCollectionRuleAgentSettings().withLogs(Arrays.asList(
                    new AgentSetting().withName(KnownAgentSettingName.MAX_DISK_QUOTA_IN_MB).withValue("5000"),
                    new AgentSetting().withName(KnownAgentSettingName.USE_TIME_RECEIVED_FOR_FORWARDED_EVENTS).withValue(
                        "1"),
                    new AgentSetting().withName(KnownAgentSettingName.TAGS).withValue("Azure, Monitoring")))),
            com.azure.core.util.Context.NONE);
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
