import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicy;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyClearKeyConfiguration;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOpenRestriction;
import com.azure.resourcemanager.mediaservices.models.ContentKeyPolicyOption;
import java.util.Arrays;

/** Samples for ContentKeyPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-update.json
     */
    /**
     * Sample code: Update a Content Key Policy.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAContentKeyPolicy(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        ContentKeyPolicy resource =
            manager
                .contentKeyPolicies()
                .getWithResponse("contoso", "contosomedia", "PolicyWithClearKeyOptionAndTokenRestriction", Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("Updated Policy")
            .withOptions(
                Arrays
                    .asList(
                        new ContentKeyPolicyOption()
                            .withName("ClearKeyOption")
                            .withConfiguration(new ContentKeyPolicyClearKeyConfiguration())
                            .withRestriction(new ContentKeyPolicyOpenRestriction())))
            .apply();
    }
}
