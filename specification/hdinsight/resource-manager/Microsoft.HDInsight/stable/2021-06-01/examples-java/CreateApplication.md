Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.hdinsight.models.ApplicationGetHttpsEndpoint;
import com.azure.resourcemanager.hdinsight.models.ApplicationProperties;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.HardwareProfile;
import com.azure.resourcemanager.hdinsight.models.Role;
import com.azure.resourcemanager.hdinsight.models.RuntimeScriptAction;
import java.util.Arrays;

/** Samples for Applications Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateApplication.json
     */
    /**
     * Sample code: Create Application.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void createApplication(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .applications()
            .define("hue")
            .withExistingCluster("rg1", "cluster1")
            .withProperties(
                new ApplicationProperties()
                    .withComputeProfile(
                        new ComputeProfile()
                            .withRoles(
                                Arrays
                                    .asList(
                                        new Role()
                                            .withName("edgenode")
                                            .withTargetInstanceCount(1)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("Standard_D12_v2")))))
                    .withInstallScriptActions(
                        Arrays
                            .asList(
                                new RuntimeScriptAction()
                                    .withName("app-install-app1")
                                    .withUri("https://.../install.sh")
                                    .withParameters("-version latest -port 20000")
                                    .withRoles(Arrays.asList("edgenode"))))
                    .withUninstallScriptActions(Arrays.asList())
                    .withHttpsEndpoints(
                        Arrays
                            .asList(
                                new ApplicationGetHttpsEndpoint()
                                    .withAccessModes(Arrays.asList("WebPage"))
                                    .withDestinationPort(20000)
                                    .withSubDomainSuffix("dss")))
                    .withApplicationType("CustomApplication")
                    .withErrors(Arrays.asList()))
            .create();
    }
}
```
