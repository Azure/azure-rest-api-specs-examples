
import com.azure.resourcemanager.horizondb.fluent.models.MicrosoftEntraAuthenticationInner;
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationProperties;

/**
 * Samples for MicrosoftEntraAuthentications CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAuthentications_CreateOrReplaceDisabled.json
     */
    /**
     * Sample code: Disable Microsoft Entra authentication declaratively.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        disableMicrosoftEntraAuthenticationDeclaratively(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAuthentications().createOrReplace("example-resource-group", "example-cluster",
            new MicrosoftEntraAuthenticationInner()
                .withProperties(new MicrosoftEntraAuthenticationProperties().withState(AuthenticationState.DISABLED)),
            com.azure.core.util.Context.NONE);
    }
}
