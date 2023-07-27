import com.azure.resourcemanager.batch.models.AutoScaleSettings;
import com.azure.resourcemanager.batch.models.BatchPoolIdentity;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.PoolIdentityType;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.UserAssignedIdentities;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_UserAssignedIdentities.json
     */
    /**
     * Sample code: CreatePool - UserAssignedIdentities.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolUserAssignedIdentities(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withIdentity(
                new BatchPoolIdentity()
                    .withType(PoolIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserAssignedIdentities(),
                            "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2",
                            new UserAssignedIdentities())))
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withVirtualMachineConfiguration(
                        new VirtualMachineConfiguration()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("Canonical")
                                    .withOffer("UbuntuServer")
                                    .withSku("18.04-LTS")
                                    .withVersion("latest"))
                            .withNodeAgentSkuId("batch.node.ubuntu 18.04")))
            .withScaleSettings(
                new ScaleSettings()
                    .withAutoScale(
                        new AutoScaleSettings()
                            .withFormula("$TargetDedicatedNodes=1")
                            .withEvaluationInterval(Duration.parse("PT5M"))))
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
