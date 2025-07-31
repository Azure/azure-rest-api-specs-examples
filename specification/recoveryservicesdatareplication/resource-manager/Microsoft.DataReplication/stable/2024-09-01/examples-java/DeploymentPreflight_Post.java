
import com.azure.resourcemanager.recoveryservicesdatareplication.fluent.models.DeploymentPreflightModelInner;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.DeploymentPreflightResource;
import java.util.Arrays;

/**
 * Samples for DeploymentPreflight Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/DeploymentPreflight_Post.json
     */
    /**
     * Sample code: Performs resource deployment validation.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void performsResourceDeploymentValidation(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.deploymentPreflights().postWithResponse("rgswagger_2024-09-01", "lnfcwsmlowbwkndkztzvaj",
            new DeploymentPreflightModelInner().withResources(Arrays
                .asList(new DeploymentPreflightResource().withName("xtgugoflfc").withType("nsnaptduolqcxsikrewvgjbxqpt")
                    .withLocation("cbsgtxkjdzwbyp").withApiVersion("otihymhvzblycdoxo"))),
            com.azure.core.util.Context.NONE);
    }
}
