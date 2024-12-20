
import com.azure.resourcemanager.automation.models.AdvancedSchedule;
import com.azure.resourcemanager.automation.models.AzureQueryProperties;
import com.azure.resourcemanager.automation.models.NonAzureQueryProperties;
import com.azure.resourcemanager.automation.models.OperatingSystemType;
import com.azure.resourcemanager.automation.models.ScheduleFrequency;
import com.azure.resourcemanager.automation.models.SoftwareUpdateConfigurationTasks;
import com.azure.resourcemanager.automation.models.SucScheduleProperties;
import com.azure.resourcemanager.automation.models.TagOperators;
import com.azure.resourcemanager.automation.models.TagSettingsProperties;
import com.azure.resourcemanager.automation.models.TargetProperties;
import com.azure.resourcemanager.automation.models.TaskProperties;
import com.azure.resourcemanager.automation.models.UpdateConfiguration;
import com.azure.resourcemanager.automation.models.WindowsProperties;
import com.azure.resourcemanager.automation.models.WindowsUpdateClasses;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SoftwareUpdateConfigurations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/
     * softwareUpdateConfiguration/createSoftwareUpdateConfiguration.json
     */
    /**
     * Sample code: Create software update configuration.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        createSoftwareUpdateConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.softwareUpdateConfigurations().define("testpatch").withExistingAutomationAccount("mygroup", "myaccount")
            .withUpdateConfiguration(new UpdateConfiguration().withOperatingSystem(OperatingSystemType.WINDOWS)
                .withWindows(new WindowsProperties().withIncludedUpdateClassifications(WindowsUpdateClasses.CRITICAL)
                    .withExcludedKbNumbers(Arrays.asList("168934", "168973")).withRebootSetting("IfRequired"))
                .withDuration(Duration.parse("PT2H0M"))
                .withAzureVirtualMachines(Arrays.asList(
                    "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01",
                    "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02",
                    "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03"))
                .withNonAzureComputerNames(Arrays.asList("box1.contoso.com", "box2.contoso.com"))
                .withTargets(new TargetProperties()
                    .withAzureQueries(Arrays.asList(new AzureQueryProperties()
                        .withScope(Arrays.asList(
                            "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources",
                            "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067"))
                        .withLocations(Arrays.asList("Japan East", "UK South"))
                        .withTagSettings(new TagSettingsProperties()
                            .withTags(mapOf("tag1", Arrays.asList("tag1Value1", "tag1Value2", "tag1Value3"), "tag2",
                                Arrays.asList("tag2Value1", "tag2Value2", "tag2Value3")))
                            .withFilterOperator(TagOperators.ALL))))
                    .withNonAzureQueries(Arrays.asList(
                        new NonAzureQueryProperties().withFunctionAlias("SavedSearch1").withWorkspaceId("WorkspaceId1"),
                        new NonAzureQueryProperties().withFunctionAlias("SavedSearch2")
                            .withWorkspaceId("WorkspaceId2")))))
            .withScheduleInfo(
                new SucScheduleProperties().withStartTime(OffsetDateTime.parse("2017-10-19T12:22:57+00:00"))
                    .withExpiryTime(OffsetDateTime.parse("2018-11-09T11:22:57+00:00")).withInterval(1L)
                    .withFrequency(ScheduleFrequency.HOUR).withTimeZone("America/Los_Angeles")
                    .withAdvancedSchedule(new AdvancedSchedule().withWeekDays(Arrays.asList("Monday", "Thursday"))))
            .withTasks(new SoftwareUpdateConfigurationTasks()
                .withPreTask(
                    new TaskProperties().withParameters(mapOf("COMPUTERNAME", "Computer1")).withSource("HelloWorld"))
                .withPostTask(new TaskProperties().withSource("GetCache")))
            .create();
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
