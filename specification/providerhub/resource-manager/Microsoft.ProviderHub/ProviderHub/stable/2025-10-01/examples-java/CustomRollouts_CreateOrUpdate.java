
import com.azure.resourcemanager.providerhub.models.CheckinManifestParams;
import com.azure.resourcemanager.providerhub.models.CustomRolloutProperties;
import com.azure.resourcemanager.providerhub.models.CustomRolloutPropertiesSpecification;
import com.azure.resourcemanager.providerhub.models.CustomRolloutSpecificationAutoProvisionConfig;
import com.azure.resourcemanager.providerhub.models.CustomRolloutSpecificationCanary;
import com.azure.resourcemanager.providerhub.models.ManifestCheckinOption;
import com.azure.resourcemanager.providerhub.models.ManifestCheckinSpecification;
import java.util.Arrays;

/**
 * Samples for CustomRollouts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/CustomRollouts_CreateOrUpdate.json
     */
    /**
     * Sample code: CustomRollouts_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void customRolloutsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.customRollouts().define("brazilUsShoeBoxTesting").withExistingProviderRegistration("Microsoft.Contoso")
            .withProperties(new CustomRolloutProperties().withSpecification(new CustomRolloutPropertiesSpecification()
                .withAutoProvisionConfig(
                    new CustomRolloutSpecificationAutoProvisionConfig().withStorage(true).withResourceGraph(true))
                .withCanary(new CustomRolloutSpecificationCanary().withRegions(Arrays.asList("brazilus")))
                .withRefreshSubscriptionRegistration(true).withRolloutId("Ev2RolloutIdGuid")
                .withManifestCheckinSpecification(new ManifestCheckinSpecification()
                    .withManifestCheckinOption(ManifestCheckinOption.ATTEMPT_AUTOMATIC_MANIFEST_CHECKIN)
                    .withManifestCheckinParams(new CheckinManifestParams().withEnvironment("Prod")
                        .withBaselineArmManifestLocation("EastUS2EUAP")))))
            .create();
    }
}
