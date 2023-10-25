import com.azure.resourcemanager.recoveryservicesdatareplication.fluent.models.DeploymentPreflightModelInner;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.DeploymentPreflightResource;
import java.util.Arrays;

/** Samples for ResourceProvider DeploymentPreflight. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/DeploymentPreflight.json
     */
    /**
     * Sample code: DeploymentPreflight.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deploymentPreflight(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .resourceProviders()
            .deploymentPreflightWithResponse(
                "rgrecoveryservicesdatareplication",
                "kjoiahxljomjcmvabaobumg",
                new DeploymentPreflightModelInner()
                    .withResources(
                        Arrays
                            .asList(
                                new DeploymentPreflightResource()
                                    .withName("xtgugoflfc")
                                    .withType("nsnaptduolqcxsikrewvgjbxqpt")
                                    .withLocation("cbsgtxkjdzwbyp")
                                    .withApiVersion("otihymhvzblycdoxo"))),
                com.azure.core.util.Context.NONE);
    }
}
