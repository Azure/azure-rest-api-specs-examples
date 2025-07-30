
import com.azure.resourcemanager.providerhub.models.DefaultRolloutProperties;
import com.azure.resourcemanager.providerhub.models.DefaultRolloutPropertiesSpecification;
import com.azure.resourcemanager.providerhub.models.DefaultRolloutSpecificationCanary;
import com.azure.resourcemanager.providerhub.models.DefaultRolloutSpecificationExpeditedRollout;
import com.azure.resourcemanager.providerhub.models.DefaultRolloutSpecificationRestOfTheWorldGroupTwo;
import java.time.Duration;
import java.util.Arrays;

/**
 * Samples for DefaultRollouts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * DefaultRollouts_CreateOrUpdate.json
     */
    /**
     * Sample code: DefaultRollouts_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void defaultRolloutsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.defaultRollouts().define("2020week10").withExistingProviderRegistration("Microsoft.Contoso")
            .withProperties(new DefaultRolloutProperties().withSpecification(new DefaultRolloutPropertiesSpecification()
                .withExpeditedRollout(new DefaultRolloutSpecificationExpeditedRollout().withEnabled(true))
                .withCanary(new DefaultRolloutSpecificationCanary().withSkipRegions(Arrays.asList("eastus2euap")))
                .withRestOfTheWorldGroupTwo(
                    new DefaultRolloutSpecificationRestOfTheWorldGroupTwo().withWaitDuration(Duration.parse("PT4H")))))
            .create();
    }
}
