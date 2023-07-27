import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.batch.models.AutoScaleSettings;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.NodeCommunicationMode;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.VMExtension;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.io.IOException;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
     */
    /**
     * Sample code: CreatePool - VirtualMachineConfiguration Extensions.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolVirtualMachineConfigurationExtensions(
        com.azure.resourcemanager.batch.BatchManager manager) throws IOException {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withVirtualMachineConfiguration(
                        new VirtualMachineConfiguration()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("Canonical")
                                    .withOffer("0001-com-ubuntu-server-focal")
                                    .withSku("20_04-lts"))
                            .withNodeAgentSkuId("batch.node.ubuntu 20.04")
                            .withExtensions(
                                Arrays
                                    .asList(
                                        new VMExtension()
                                            .withName("batchextension1")
                                            .withPublisher("Microsoft.Azure.KeyVault")
                                            .withType("KeyVaultForLinux")
                                            .withTypeHandlerVersion("2.0")
                                            .withAutoUpgradeMinorVersion(true)
                                            .withEnableAutomaticUpgrade(true)
                                            .withSettings(
                                                SerializerFactory
                                                    .createDefaultManagementSerializerAdapter()
                                                    .deserialize(
                                                        "{\"authenticationSettingsKey\":\"authenticationSettingsValue\",\"secretsManagementSettingsKey\":\"secretsManagementSettingsValue\"}",
                                                        Object.class,
                                                        SerializerEncoding.JSON))))))
            .withScaleSettings(
                new ScaleSettings()
                    .withAutoScale(
                        new AutoScaleSettings()
                            .withFormula("$TargetDedicatedNodes=1")
                            .withEvaluationInterval(Duration.parse("PT5M"))))
            .withTargetNodeCommunicationMode(NodeCommunicationMode.DEFAULT)
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
