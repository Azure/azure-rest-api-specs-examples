Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.machinelearning.models.ApplicationSharingPolicy;
import com.azure.resourcemanager.machinelearning.models.AssignedUser;
import com.azure.resourcemanager.machinelearning.models.ComputeInstance;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceAuthorizationType;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceProperties;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceSshSettings;
import com.azure.resourcemanager.machinelearning.models.PersonalComputeInstanceSettings;
import com.azure.resourcemanager.machinelearning.models.ResourceId;
import com.azure.resourcemanager.machinelearning.models.SshPublicAccess;
import java.util.HashMap;
import java.util.Map;

/** Samples for Compute CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/createOrUpdate/ComputeInstance.json
     */
    /**
     * Sample code: Create an ComputeInstance Compute.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createAnComputeInstanceCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .computes()
            .define("compute123")
            .withExistingWorkspace("testrg123", "workspaces123")
            .withRegion("eastus")
            .withProperties(
                new ComputeInstance()
                    .withProperties(
                        new ComputeInstanceProperties()
                            .withVmSize("STANDARD_NC6")
                            .withSubnet(new ResourceId().withId("test-subnet-resource-id"))
                            .withApplicationSharingPolicy(ApplicationSharingPolicy.PERSONAL)
                            .withSshSettings(
                                new ComputeInstanceSshSettings().withSshPublicAccess(SshPublicAccess.DISABLED))
                            .withComputeInstanceAuthorizationType(ComputeInstanceAuthorizationType.PERSONAL)
                            .withPersonalComputeInstanceSettings(
                                new PersonalComputeInstanceSettings()
                                    .withAssignedUser(
                                        new AssignedUser()
                                            .withObjectId("00000000-0000-0000-0000-000000000000")
                                            .withTenantId("00000000-0000-0000-0000-000000000000")))))
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
```
