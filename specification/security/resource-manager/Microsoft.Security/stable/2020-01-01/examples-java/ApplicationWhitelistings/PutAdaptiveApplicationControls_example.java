import com.azure.core.util.Context;
import com.azure.resourcemanager.security.models.AdaptiveApplicationControlGroup;
import com.azure.resourcemanager.security.models.ConfigurationStatus;
import com.azure.resourcemanager.security.models.EnforcementMode;
import com.azure.resourcemanager.security.models.EnforcementSupport;
import com.azure.resourcemanager.security.models.FileType;
import com.azure.resourcemanager.security.models.PathRecommendation;
import com.azure.resourcemanager.security.models.ProtectionMode;
import com.azure.resourcemanager.security.models.PublisherInfo;
import com.azure.resourcemanager.security.models.RecommendationAction;
import com.azure.resourcemanager.security.models.RecommendationType;
import com.azure.resourcemanager.security.models.UserRecommendation;
import com.azure.resourcemanager.security.models.VmRecommendation;
import java.util.Arrays;

/** Samples for AdaptiveApplicationControls Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/PutAdaptiveApplicationControls_example.json
     */
    /**
     * Sample code: Update an application control machine group by adding a new application.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void updateAnApplicationControlMachineGroupByAddingANewApplication(
        com.azure.resourcemanager.security.SecurityManager manager) {
        AdaptiveApplicationControlGroup resource =
            manager.adaptiveApplicationControls().getWithResponse("centralus", "ERELGROUP1", Context.NONE).getValue();
        resource
            .update()
            .withEnforcementMode(EnforcementMode.AUDIT)
            .withProtectionMode(
                new ProtectionMode()
                    .withExe(EnforcementMode.AUDIT)
                    .withMsi(EnforcementMode.NONE)
                    .withScript(EnforcementMode.NONE))
            .withVmRecommendations(
                Arrays
                    .asList(
                        new VmRecommendation()
                            .withConfigurationStatus(ConfigurationStatus.CONFIGURED)
                            .withRecommendationAction(RecommendationAction.RECOMMENDED)
                            .withResourceId(
                                "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/erelh-stable/providers/microsoft.compute/virtualmachines/erelh-16090")
                            .withEnforcementSupport(EnforcementSupport.SUPPORTED),
                        new VmRecommendation()
                            .withConfigurationStatus(ConfigurationStatus.CONFIGURED)
                            .withRecommendationAction(RecommendationAction.RECOMMENDED)
                            .withResourceId(
                                "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/matanvs/providers/microsoft.compute/virtualmachines/matanvs19")
                            .withEnforcementSupport(EnforcementSupport.SUPPORTED)))
            .withPathRecommendations(
                Arrays
                    .asList(
                        new PathRecommendation()
                            .withPath("[Exe] O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US\\*\\*\\0.0.0.0")
                            .withAction(RecommendationAction.RECOMMENDED)
                            .withType(RecommendationType.fromString("PublisherSignature"))
                            .withPublisherInfo(
                                new PublisherInfo()
                                    .withPublisherName("O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US")
                                    .withProductName("*")
                                    .withBinaryName("*")
                                    .withVersion("0.0.0.0"))
                            .withCommon(true)
                            .withUserSids(Arrays.asList("S-1-1-0"))
                            .withUsernames(
                                Arrays
                                    .asList(
                                        new UserRecommendation()
                                            .withUsername("Everyone")
                                            .withRecommendationAction(RecommendationAction.RECOMMENDED)))
                            .withFileType(FileType.EXE)
                            .withConfigurationStatus(ConfigurationStatus.CONFIGURED),
                        new PathRecommendation()
                            .withPath("%OSDRIVE%\\WINDOWSAZURE\\SECAGENT\\WASECAGENTPROV.EXE")
                            .withAction(RecommendationAction.RECOMMENDED)
                            .withType(RecommendationType.fromString("ProductSignature"))
                            .withPublisherInfo(
                                new PublisherInfo()
                                    .withPublisherName("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN")
                                    .withProductName("MICROSOFT® COREXT")
                                    .withBinaryName("*")
                                    .withVersion("0.0.0.0"))
                            .withCommon(true)
                            .withUserSids(Arrays.asList("S-1-1-0"))
                            .withUsernames(
                                Arrays
                                    .asList(
                                        new UserRecommendation()
                                            .withUsername("NT AUTHORITY\\SYSTEM")
                                            .withRecommendationAction(RecommendationAction.RECOMMENDED)))
                            .withFileType(FileType.EXE)
                            .withConfigurationStatus(ConfigurationStatus.CONFIGURED),
                        new PathRecommendation()
                            .withPath("%OSDRIVE%\\WINDOWSAZURE\\PACKAGES_201973_7415\\COLLECTGUESTLOGS.EXE")
                            .withAction(RecommendationAction.RECOMMENDED)
                            .withType(RecommendationType.fromString("PublisherSignature"))
                            .withPublisherInfo(
                                new PublisherInfo()
                                    .withPublisherName("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN")
                                    .withProductName("*")
                                    .withBinaryName("*")
                                    .withVersion("0.0.0.0"))
                            .withCommon(true)
                            .withUserSids(Arrays.asList("S-1-1-0"))
                            .withUsernames(
                                Arrays
                                    .asList(
                                        new UserRecommendation()
                                            .withUsername("NT AUTHORITY\\SYSTEM")
                                            .withRecommendationAction(RecommendationAction.RECOMMENDED)))
                            .withFileType(FileType.EXE)
                            .withConfigurationStatus(ConfigurationStatus.CONFIGURED),
                        new PathRecommendation()
                            .withPath("C:\\directory\\file.exe")
                            .withAction(RecommendationAction.ADD)
                            .withType(RecommendationType.fromString("File"))
                            .withCommon(true)))
            .apply();
    }
}
