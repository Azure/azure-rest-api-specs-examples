
import com.azure.resourcemanager.machinelearning.models.ApplicationSharingPolicy;
import com.azure.resourcemanager.machinelearning.models.AssignedUser;
import com.azure.resourcemanager.machinelearning.models.ComputeInstance;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceAuthorizationType;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceProperties;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceSshSettings;
import com.azure.resourcemanager.machinelearning.models.ComputePowerAction;
import com.azure.resourcemanager.machinelearning.models.ComputeSchedules;
import com.azure.resourcemanager.machinelearning.models.ComputeStartStopSchedule;
import com.azure.resourcemanager.machinelearning.models.ComputeTriggerType;
import com.azure.resourcemanager.machinelearning.models.Cron;
import com.azure.resourcemanager.machinelearning.models.PersonalComputeInstanceSettings;
import com.azure.resourcemanager.machinelearning.models.ScheduleStatus;
import com.azure.resourcemanager.machinelearning.models.SshPublicAccess;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Compute CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/createOrUpdate/ComputeInstanceWithSchedules.json
     */
    /**
     * Sample code: Create an ComputeInstance Compute with Schedules.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createAnComputeInstanceComputeWithSchedules(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().define("compute123").withExistingWorkspace("testrg123", "workspaces123").withRegion("eastus")
            .withProperties(new ComputeInstance().withProperties(new ComputeInstanceProperties()
                .withVmSize("STANDARD_NC6").withApplicationSharingPolicy(ApplicationSharingPolicy.PERSONAL)
                .withSshSettings(new ComputeInstanceSshSettings().withSshPublicAccess(SshPublicAccess.DISABLED))
                .withComputeInstanceAuthorizationType(ComputeInstanceAuthorizationType.PERSONAL)
                .withPersonalComputeInstanceSettings(new PersonalComputeInstanceSettings()
                    .withAssignedUser(new AssignedUser().withObjectId("00000000-0000-0000-0000-000000000000")
                        .withTenantId("00000000-0000-0000-0000-000000000000")))
                .withSchedules(
                    new ComputeSchedules()
                        .withComputeStartStop(
                            Arrays
                                .asList(new ComputeStartStopSchedule().withStatus(ScheduleStatus.ENABLED)
                                    .withAction(ComputePowerAction.STOP).withTriggerType(ComputeTriggerType.CRON)
                                    .withCron(new Cron().withStartTime("2021-04-23T01:30:00")
                                        .withTimeZone("Pacific Standard Time").withExpression("0 18 * * *")))))))
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
