<script>
    export let title = "A card with media.";
    export let description = "It's all in this card. It's a veritable smorgasbord of card features.";
    export let image = '';
    export let likeCount = 0;
    export let commentCount = 0;
    import Card, {Content, PrimaryAction, Media, Actions, ActionButtons, ActionIcons} from '@smui/card';
    import IconButton from '@smui/icon-button';
    let clicked = 0;
</script>

<style>
    .card-section {
        margin-bottom: 20px;
        width: 250px;
        margin-left: auto;
        margin-right: auto;
    }
  h2 {
      font-weight: normal;
  }
</style>
<div class="card-section">
    <Card>
        <PrimaryAction on:click={() => clicked++}>
            <Media style="background-image:url({image})" class="card-media-16x9" aspectRatio="16x9"></Media>
            <Content class="mdc-typography--body2">
            <h2 class="mdc-typography--headline6" style="margin: 0;">{title}</h2>
            <p>{description}</p>
            </Content>
        </PrimaryAction>
        <Actions>
            <ActionIcons>
                <div style="display: flex; align-items: center;">
                    <IconButton class="material-icons" on:click={() => clicked++} title="Like">thumb_up</IconButton> {likeCount}
                </div>
                <div style="display: flex; align-items: center;">
                    <IconButton class="material-icons" on:click={() => clicked++} title="Comment">comment</IconButton> {commentCount}
                </div>
            </ActionIcons>
        </Actions>
    </Card>
</div>