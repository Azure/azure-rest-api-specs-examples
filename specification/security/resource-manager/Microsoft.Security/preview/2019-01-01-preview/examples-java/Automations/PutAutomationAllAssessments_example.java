import com.azure.resourcemanager.security.models.AutomationActionLogicApp;
import com.azure.resourcemanager.security.models.AutomationScope;
import com.azure.resourcemanager.security.models.AutomationSource;
import com.azure.resourcemanager.security.models.EventSource;
import java.io.IOException;
import java.net.URL;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Automations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationAllAssessments_example.json
     */
    /**
     * Sample code: Create or update a security automation for all assessments (including all severities).
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateASecurityAutomationForAllAssessmentsIncludingAllSeverities(
        com.azure.resourcemanager.security.SecurityManager manager) throws IOException {
        manager
            .automations()
            .define("exampleAutomation")
            .withRegion("Central US")
            .withExistingResourceGroup("exampleResourceGroup")
            .withTags(mapOf())
            .withDescription(
                "An example of a security automation that triggers one LogicApp resource (myTest1) on any security"
                    + " assessment")
            .withIsEnabled(true)
            .withScopes(
                Arrays
                    .asList(
                        new AutomationScope()
                            .withDescription(
                                "A description that helps to identify this scope - for example: security assessments"
                                    + " that relate to the resource group myResourceGroup within the subscription"
                                    + " a5caac9c-5c04-49af-b3d0-e204f40345d5")
                            .withScopePath(
                                "/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup")))
            .withSources(Arrays.asList(new AutomationSource().withEventSource(EventSource.ASSESSMENTS)))
            .withActions(
                Arrays
                    .asList(
                        new AutomationActionLogicApp()
                            .withLogicAppResourceId(
                                "/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1")
                            .withUri(new URL("https://exampleTriggerUri1.com"))))
            .create();
    }

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
