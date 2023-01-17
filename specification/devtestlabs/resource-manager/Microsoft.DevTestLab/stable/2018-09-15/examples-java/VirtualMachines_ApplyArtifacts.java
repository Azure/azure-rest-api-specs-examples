import com.azure.resourcemanager.devtestlabs.models.ApplyArtifactsRequest;
import com.azure.resourcemanager.devtestlabs.models.ArtifactInstallProperties;
import java.util.Arrays;

/** Samples for VirtualMachines ApplyArtifacts. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_ApplyArtifacts.json
     */
    /**
     * Sample code: VirtualMachines_ApplyArtifacts.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesApplyArtifacts(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .applyArtifacts(
                "resourceGroupName",
                "{labName}",
                "{vmName}",
                new ApplyArtifactsRequest()
                    .withArtifacts(
                        Arrays
                            .asList(
                                new ArtifactInstallProperties()
                                    .withArtifactId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/public"
                                            + " repo/artifacts/windows-restart"))),
                com.azure.core.util.Context.NONE);
    }
}
