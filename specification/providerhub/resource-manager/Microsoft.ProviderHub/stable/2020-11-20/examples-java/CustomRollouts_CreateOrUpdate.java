
import com.azure.resourcemanager.providerhub.models.CustomRolloutProperties;
import com.azure.resourcemanager.providerhub.models.CustomRolloutPropertiesSpecification;
import com.azure.resourcemanager.providerhub.models.CustomRolloutSpecificationCanary;
import java.util.Arrays;

/**
 * Samples for CustomRollouts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/
     * CustomRollouts_CreateOrUpdate.json
     */
    /**
     * Sample code: CustomRollouts_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void customRolloutsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.customRollouts().define("brazilUsShoeBoxTesting").withExistingProviderRegistration("Microsoft.Contoso")
            .withProperties(new CustomRolloutProperties().withSpecification(new CustomRolloutPropertiesSpecification()
                .withCanary(new CustomRolloutSpecificationCanary().withRegions(Arrays.asList("brazilus")))))
            .create();
    }
}
