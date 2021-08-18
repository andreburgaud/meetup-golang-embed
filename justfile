clean:
    @for dir in `ls -1d ?-*`; do just $dir/clean; done

build:
    @for dir in `ls -1d ?-*`; do just $dir/build; done